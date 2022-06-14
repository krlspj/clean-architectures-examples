#ifndef MyController_hpp
#define MyController_hpp

#include <memory>
#include <iostream>
#include "dto/DTOs.hpp"

#include "oatpp/web/server/api/ApiController.hpp"
#include "oatpp/core/macro/codegen.hpp"
#include "oatpp/core/macro/component.hpp"

#include "../../core/ports/UserUseCase.cpp"

#include OATPP_CODEGEN_BEGIN(ApiController) //<-- Begin Codegen

/**
 * Sample Api Controller.
 */
class UserController : public oatpp::web::server::api::ApiController {
private:
    std::shared_ptr<UserUseCase> userUseCase;

public:
  /**
   * Constructor with object mapper.
   * @param objectMapper - default object mapper used to serialize/deserialize DTOs.
   */
  UserController(OATPP_COMPONENT(std::shared_ptr<ObjectMapper>, objectMapper))
    : oatpp::web::server::api::ApiController(objectMapper)
  {}

  void setUserUseCase(std::shared_ptr<UserUseCase> useCase) {
    userUseCase = useCase;
  }

public:

  ENDPOINT("GET", "/", root) {
    return createResponse(Status::CODE_200, "");
  }
  
  ENDPOINT("GET", "/users/{userId}", getUserById, PATH(String, userId)) {
    OATPP_LOGD("getUserById", "userId=%s", userId->std_str().c_str());

    long lUserId = strtol(userId->std_str().c_str(),NULL,10);

    auto user = userUseCase->getUserById(lUserId);
    
    auto dto = UserDto::createShared();
    
    dto->id = user->id();
    dto->name = user->name().c_str();
    dto->email = user->email().c_str();
    dto->statusCode = 200;

    return createDtoResponse(Status::CODE_200, dto);
  }
  
};

#include OATPP_CODEGEN_END(ApiController) //<-- End Codegen

#endif /* UserController_hpp */