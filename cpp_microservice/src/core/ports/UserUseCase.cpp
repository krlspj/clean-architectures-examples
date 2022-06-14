#ifndef USERUSECASE_CPP
#define USERUSECASE_CPP
#include <memory>
#include "../domain/User.cpp"

class UserUseCase {
public:
    virtual std::shared_ptr<User> getUserById(long id) = 0;
};
#endif // USERUSECASE_CPP