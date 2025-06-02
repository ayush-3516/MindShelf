FROM nginx:alpine

# Copy custom nginx config
COPY nginx.conf /etc/nginx/nginx.conf

# Copy static files
# The public folder will be mounted as a volume in docker-compose.yml

# Expose ports
EXPOSE 80
EXPOSE 443

# Command to run
CMD ["nginx", "-g", "daemon off;"]