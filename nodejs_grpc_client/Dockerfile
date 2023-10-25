# Use an official Node.js runtime as a parent image
FROM node:18-alpine

# Set the working directory in the container
WORKDIR /app

# Install Python and build tools
RUN apk add --no-cache python3 make g++ gcc

# Copy the package.json and package-lock.json to the container
COPY package*.json ./

# Install application dependencies
RUN npm install

# Copy the rest of your application's source code
COPY . .

# Define the command to run your Node.js application
CMD [ "node", "server.js" ]

