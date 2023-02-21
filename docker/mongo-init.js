db = db.getSiblingDB('gin_mongo');
db.createUser(
    {
        user: "mongo",
        pwd: "mongo",
        roles: [
            {
                role: "readWrite",
                db: "mongo"
            }
        ]
    }
)
