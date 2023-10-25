# Use an appropriate Python base image
FROM python:3.9-slim

# Set the working directory
WORKDIR /app

# Create a subdirectory for the proto folder
RUN mkdir proto

# Copy your Python script, requirements file, and the proto files to the container
COPY server.py requirements.txt /app/
COPY proto/* /app/proto/

# Install dependencies from requirements.txt
RUN pip install -r requirements.txt

# Specify the command to run your Python script
CMD ["python", "server.py"]

