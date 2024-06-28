**CRUD API’s and Authentication endpoints**
**CRUD API’s for user details :**


1. /createUserDetails -POST method - to create user details with JSON data of id, name, place, age.

2. /getAllUserDetails - GET method - to get all user details.

3. getUserDetailsById/:id - GET method - to get specific user detail using id parameter in request.

4. /updateUserDetailsById/:id - PUT method - to update user detail of specific user using id parameter in request.

5. /deleteUserDetailsById/:id - DELETE method - to delete user detail of specific user using id parameter in request.

**AUTHENTICATION Endpoints:**

1. /createAuthUser - POST method - to create authenticate username and password(password encrypted with bcrypt)

2. /authUserLogin - POST method - to authenticate username and password is valid or not(password decrypt by bcrypt) and return jwt token

3. /getAuthUser - GET method - to access a Auth user data using bearer token in authentication header(exp time of token is 2 minutes)


JSON body of user details:
 Example:   {
        "id": "90",
        "name": "Prabhu",
        "place": "pollachi",
        "age": 23
    }

JSON body of Auth user:
Example: {
    "username":"gopal",
    "password":"stringOfGopal"
}

All the entire data stored in mongo DB with two different collection(userDetails and AuthenticationUserDetails) in database(data-gin as db name) mongoDB default port.

---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------

**Steps to run server locally**

1. Install Go(preferred latest version).
2. Install mongod and setup mongoDB compass(GUI)
3. Connect default port of mongDB  in mongoDB compass
4. I setted default port 8008 to run server in config/env.go you can change in that env.go file(PORT) check any other server or something running on your local machine in that port and change it accordingly.
5. Clone this go-gin-crud repo to your machine.
6. Go to go-gin-crud directory and run the command(go run main.go).
