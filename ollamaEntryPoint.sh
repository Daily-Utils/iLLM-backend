#!/bin/sh

curl -sSL https://ollama.com/install.sh | sh

# Start the ollama server in the background and redirect output to a log file
# ollama serve > /app/ollama_serve.log 2>&1 &
ollama serve


# Wait for the server to start
# sleep 20

# run_model() {
#   local retries=5
#   local count=0
#   echo "Im hereeee"
#   until ollama run llama3 || [ $count -eq $retries ]; do
#     count=$((count + 1))
#     echo "Retrying to run the model... Attempt $count/$retries"
#     sleep 10
#   done

#   if [ $count -eq $retries ]; then
#     echo "Failed to run the model after $retries attempts."
#     exit 1
#   fi
# }

# Run the model with retries in a child process
# run_model