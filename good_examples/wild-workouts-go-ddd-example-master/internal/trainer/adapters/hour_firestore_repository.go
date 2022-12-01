package adapters

import (
	"context"
	"time"

	"cloud.google.com/go/firestore"
	"github.com/ThreeDotsLabs/wild-workouts-go-ddd-example/internal/trainer/domain/hour"
	"github.com/pkg/errors"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type FirestoreHourRepository struct {
	firestoreClient *firestore.Client
	hourFactory     hour.Factory
}

func NewFirestoreHourRepository(firestoreClient *firestore.Client, hourFactory hour.Factory) *FirestoreHourRepository {
	if firestoreClient == nil {
		panic("missing firestoreClient")
	}
	if hourFactory.IsZero() {
		panic("missing hourFactory")
	}

	return &FirestoreHourRepository{firestoreClient, hourFactory}
}

func (f FirestoreHourRepository) GetHour(ctx context.Context, time time.Time) (*hour.Hour, error) {
	date, err := f.getDateDTO(
		// getDateDTO should be used both for transactional and non transactional query,
		// the best way for that is to use closure
		func() (doc *firestore.DocumentSnapshot, err error) {
			return f.documentRef(time).Get(ctx)
		},
		time,
	)
	if err != nil {
		return nil, err
	}

	hourFromDb, err := f.domainHourFromDateDTO(date, time)
	if err != nil {
		return nil, err
	}

	return hourFromDb, err
}

func (f FirestoreHourRepository) UpdateHour(
	ctx context.Context,
	hourTime time.Time,
	updateFn func(h *hour.Hour) (*hour.Hour, error),
) error {
	err := f.firestoreClient.RunTransaction(ctx, func(ctx context.Context, transaction *firestore.Transaction) error {
		dateDocRef := f.documentRef(hourTime)

		firebaseDate, err := f.getDateDTO(
			// getDateDTO should be used both for transactional and non transactional query,
			// the best way for that is to use closure
			func() (doc *firestore.DocumentSnapshot, err error) {
				return transaction.Get(dateDocRef)
			},
			hourTime,
		)
		if err != nil {
			return err
		}

		hourFromDB, err := f.domainHourFromDateDTO(firebaseDate, hourTime)
		if err != nil {
			return err
		}

		updatedHour, err := updateFn(hourFromDB)
		if err != nil {
			return errors.Wrap(err, "unable to update hour")
		}
		updateHourInDataDTO(updatedHour, &firebaseDate)

		return transaction.Set(dateDocRef, firebaseDate)
	})

	return errors.Wrap(err, "firestore transaction failed")
}

func (f FirestoreHourRepository) trainerHoursCollection() *firestore.CollectionRef {
	return f.firestoreClient.Collection("trainer-hours")
}

func (f FirestoreHourRepository) documentRef(hourTime time.Time) *firestore.DocumentRef {
	return f.trainerHoursCollection().Doc(hourTime.Format("2006-01-02"))
}

func (f FirestoreHourRepository) getDateDTO(
	getDocumentFn func() (doc *firestore.DocumentSnapshot, err error),
	dateTime time.Time,
) (DateModel, error) {
	doc, err := getDocumentFn()
	if status.Code(err) == codes.NotFound {
		// in reality this date exists, even if it's not persisted
		return NewEmptyDateDTO(dateTime), nil
	}
	if err != nil {
		return DateModel{}, err
	}

	date := DateModel{}
	if err := doc.DataTo(&date); err != nil {
		return DateModel{}, errors.Wrap(err, "unable to unmarshal DateModel from Firestore")
	}

	return date, nil
}

// for now we are keeping backward comparability, because of that it's a bit messy and overcomplicated
// todo - we will clean it up later with CQRS :-)
func (f FirestoreHourRepository) domainHourFromDateDTO(date DateModel, hourTime time.Time) (*hour.Hour, error) {
	firebaseHour, found := findHourInDateDTO(date, hourTime)
	if !found {
		// in reality this date exists, even if it's not persisted
		return f.hourFactory.NewNotAvailableHour(hourTime)
	}

	availability, err := mapAvailabilityFromDTO(firebaseHour)
	if err != nil {
		return nil, err
	}

	return f.hourFactory.UnmarshalHourFromDatabase(firebaseHour.Hour.Local(), availability)
}

// for now we are keeping backward comparability, because of that it's a bit messy and overcomplicated
// todo - we will clean it up later with CQRS :-)
func updateHourInDataDTO(updatedHour *hour.Hour, firebaseDate *DateModel) {
	firebaseHourDTO := domainHourToDTO(updatedHour)

	hourFound := false
	for i := range firebaseDate.Hours {
		if !firebaseDate.Hours[i].Hour.Equal(updatedHour.Time()) {
			continue
		}

		firebaseDate.Hours[i] = firebaseHourDTO
		hourFound = true
		break
	}

	if !hourFound {
		firebaseDate.Hours = append(firebaseDate.Hours, firebaseHourDTO)
	}

	firebaseDate.HasFreeHours = false
	for _, h := range firebaseDate.Hours {
		if h.Available {
			firebaseDate.HasFreeHours = true
			break
		}
	}
}

func mapAvailabilityFromDTO(firebaseHour HourModel) (hour.Availability, error) {
	if firebaseHour.Available && !firebaseHour.HasTrainingScheduled {
		return hour.Available, nil
	}
	if !firebaseHour.Available && firebaseHour.HasTrainingScheduled {
		return hour.TrainingScheduled, nil
	}
	if !firebaseHour.Available && !firebaseHour.HasTrainingScheduled {
		return hour.NotAvailable, nil
	}

	return hour.Availability{}, errors.Errorf(
		"unsupported values - Available: %t, HasTrainingScheduled: %t",
		firebaseHour.Available,
		firebaseHour.HasTrainingScheduled,
	)
}

func domainHourToDTO(updatedHour *hour.Hour) HourModel {
	return HourModel{
		Available:            updatedHour.IsAvailable(),
		HasTrainingScheduled: updatedHour.HasTrainingScheduled(),
		Hour:                 updatedHour.Time(),
	}
}

func findHourInDateDTO(firebaseDate DateModel, time time.Time) (HourModel, bool) {
	for i := range firebaseDate.Hours {
		firebaseHour := firebaseDate.Hours[i]

		if !firebaseHour.Hour.Equal(time) {
			continue
		}

		return firebaseHour, true
	}

	return HourModel{}, false
}

func NewEmptyDateDTO(t time.Time) DateModel {
	return DateModel{
		Date: t.UTC().Truncate(time.Hour * 24),
	}
}
