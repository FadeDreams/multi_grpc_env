import sys
import grpc
sys.path.insert(0, './proto/')
import user_pb2
import user_pb2_grpc


# sys.path.append('/proto/')

# Define the gRPC server address
# server_address = 'localhost:8080'
server_address = 'go-grpc-server:8080'


# Create a gRPC channel to connect to the server
channel = grpc.insecure_channel(server_address)

# Create a gRPC stub for the UserService
stub = user_pb2_grpc.UserServiceStub(channel)

# Create a request message
request = user_pb2.UserRequest(id='1')
# Make the gRPC call to the GetUser method
response = stub.GetUser(request)


# Create a request for GetAllUsers (no request message needed for an Empty request)
request2 = user_pb2.Empty()
# Make the gRPC call to the GetAllUsers method
response2 = stub.GetAllUsers(request)

# Handle the response
if response2:
    for user in response2.users:
        print(f"User ID: {user.id}, Name: {user.name}, Location: {user.location}, Title: {user.title}")
else:
    print("Error calling GetAllUsers")

# Handle the response
if response:
    print(f"User ID: {response.id}, Name: {response.name}, Location: {response.location}, Title: {response.title}")
else:
    print("Error calling GetUser")

