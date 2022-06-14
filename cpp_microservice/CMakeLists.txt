cmake_minimum_required(VERSION 3.11)

set(project_name cpp-microservice) 

project(${project_name})

set(CMAKE_CXX_STANDARD 17)

add_library(${project_name}-lib
        src/AppComponent.hpp
	#src/ServiceLocator.hpp
	#src/SLModules.cpp
	src/core/ports/UserRepository.cpp
	src/adapters/repositories/UserRepositoryImpl.cpp
        src/adapters/controllers/UserController.cpp
        src/adapters/controllers/UserController.hpp
        src/adapters/controllers/dto/DTOs.hpp
	src/core/services/UserService.cpp
)

## link libs

find_package(oatpp 1.2.5 REQUIRED)

target_link_libraries(${project_name}-lib
        PUBLIC oatpp::oatpp
        PUBLIC oatpp::oatpp-test
)

target_include_directories(${project_name}-lib PUBLIC src)

## add executables

add_executable(${project_name}-exe
        src/App.cpp
	)
target_link_libraries(${project_name}-exe ${project_name}-lib)
add_dependencies(${project_name}-exe ${project_name}-lib)

set_target_properties(${project_name}-lib ${project_name}-exe PROPERTIES
        CXX_STANDARD 17
        CXX_EXTENSIONS OFF
        CXX_STANDARD_REQUIRED ON
)