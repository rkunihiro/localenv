// https://docs.mongodb.com/manual/reference/method/db.createUser/
db.createUser({
    user: "username",
    pwd: "password",
    roles: [
        {
            role: "readWrite",
            db: "test",
        },
    ],
});

db.createCollection("books");
db.books.insertMany([
    {
        isbn: "4873112699",
        title: "GNU Make",
        author: "Robert Mecklenburg",
        publisher: "O'REILLY",
    },
]);
