#ifndef USER_CPP
#define USER_CPP
#include <string>

using namespace std;

class User {
private:
    long _id;
    string _name;
    string _email;
public:
    User();
    virtual ~User();

    const long& id() const { return _id; } 
    void id(const long& id) { _id = id; } 

    const std::string& name() const { return _name; } 
    void name(const std::string& name) { _name = name; } 

    const std::string& email() const { return _email; } 
    void email(const std::string& email) { _email = email; } 
        
};

User::User() { }

User::~User() { }

#endif // USER_CPP