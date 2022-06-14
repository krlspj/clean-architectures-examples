#include "../ports/UserUseCase.cpp"
#include "../ports/UserRepository.cpp"
#include "../domain/User.cpp"

class UserService : public UserUseCase {
private:
    std::shared_ptr<UserRepository> userRepository;
    
public:
    UserService(std::shared_ptr<UserRepository> repo);
    virtual ~UserService();

    std::shared_ptr<User> getUserById(long id) override;
};


UserService::UserService(std::shared_ptr<UserRepository> repo) : userRepository(repo){ }

UserService::~UserService(){ }

std::shared_ptr<User> UserService::getUserById(long id) {
    std::shared_ptr<User> user = userRepository->findById(id);

    return user;
}