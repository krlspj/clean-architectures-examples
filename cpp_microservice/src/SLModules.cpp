#include "./ServiceLocator.hpp"
#include "./core/ports/UserRepository.cpp"
#include "./adapters/repositories/UserRepositoryImpl.cpp"
#include "./core/services/UserService.cpp"
#include "./core/ports/UserUseCase.cpp"

class RepositoriesSLModule : public ServiceLocator::Module {
public:
  void load() override {
    bind<UserRepository>("UserService").to<UserRepositoryImpl>([](SLContext_sptr slc) {
      return new UserRepositoryImpl();
    });
  }
};

// ServicesSLModule is intimate with Bar, it knows what dependencies Bar has ..
class ServicesSLModule : public ServiceLocator::Module {
public:
    void load() override {
		bind<UserUseCase>("UserService").to<UserService>([] (SLContext_sptr slc) { 
			return new UserService(slc->resolve<UserRepository>("UserService"));
		});
	}
};