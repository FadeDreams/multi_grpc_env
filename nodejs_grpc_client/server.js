const grpc = require('grpc');
const path = require('path');
const protoLoader = require('@grpc/proto-loader');

const PROTO_PATH = path.resolve(__dirname, 'proto/user.proto'); // Replace with the actual path to your .proto file

const packageDefinition = protoLoader.loadSync(PROTO_PATH, {
    keepCase: true,
    longs: String,
    enums: String,
    defaults: true,
    oneofs: true,
});

const proto = grpc.loadPackageDefinition(packageDefinition);
//const client = new proto.user.UserService('localhost:8080', grpc.credentials.createInsecure());
const serverAddress = process.env.SERVER_ADDRESS || 'go-grpc-server:8080';
const client = new proto.user.UserService(serverAddress, grpc.credentials.createInsecure());


const request = { id: '1' };

client.getUser(request, (error, response) => {
    if (!error) {
        console.log(`User ID: ${response.id}, Name: ${response.name}, Location: ${response.location}, Title: ${response.title}`);
    } else {
        console.error(`Error calling GetUser: ${error}`);
    }
});


const emptyRequest = {}; // Create an empty request
client.getAllUsers(emptyRequest, (error, response) => {
    if (!error) {
        console.log('All Users:');
        for (const user of response.users) {
            console.log(`User ID: ${user.id}, Name: ${user.name}, Location: ${user.location}, Title: ${user.title}`);
        }
    } else {
        console.error(`Error calling GetAllUsers: ${error}`);
    }
});



