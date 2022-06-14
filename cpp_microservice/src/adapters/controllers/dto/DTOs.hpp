#ifndef DTOs_hpp
#define DTOs_hpp

#include "oatpp/core/macro/codegen.hpp"
#include "oatpp/core/Types.hpp"

#include OATPP_CODEGEN_BEGIN(DTO)

/**
 *  Data Transfer Object. Object containing fields only.
 *  Used in API for serialization/deserialization and validation
 */
class UserDto : public oatpp::DTO {
  
  DTO_INIT(UserDto, DTO)
  
  /* Fields Definition */
  DTO_FIELD(Int64, id);
  DTO_FIELD(String, name);
  DTO_FIELD(String, email);
  DTO_FIELD(Int32, statusCode);


  /* Swagger info */
  DTO_FIELD_INFO(name) { 
    info->description = "user full name"; 
  }
  DTO_FIELD_INFO(email) { 
    info->description = "user full email";
  }
  
};

#include OATPP_CODEGEN_END(DTO)

#endif /* DTOs_hpp */
