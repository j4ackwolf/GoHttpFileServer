# GoHttpFileServer

GoHttpFileServer (Go Http File Server) is a simple HTTP/HTTPS file server with a REST API extension.
The project is designed for use in CI/CD Pipelines and serves as a storage solution for distributing artifacts,
build outputs, releases, and other valuable files across various development projects and teams.

### Key Features:

- HTTP/HTTPS Server: GoHttpFileServer provides a robust and secure access to files and directories via 
the HTTP and HTTPS protocols, enabling seamless integration with different CI/CD tools and workflows.


- REST API: The project offers a straightforward REST API that allows automated management of
files and directories. Teams can automate processes such as uploading, downloading, renaming, and file management
through the API.


- Artifact and Build Storage: GoHttpFileServer acts as a safe and convenient repository for 
storing valuable artifacts, binary outputs, releases, and other essential data needed during the development and 
deployment of projects.


- Basic Authentication: To ensure security, the server supports basic authentication, allowing access to files and
API only to authorized users or teams.


- Content Listing: GoHttpFileServer enables users to view the content of directories and files through a web interface,
facilitating quick and efficient search for specific data.


- API File Upload: With a straightforward API using basic authentication, users can easily upload files to the server,
providing a flexible and convenient method for data transfer.


- File and Directory Renaming: GoHttpFileServer allows renaming files and directories through the API, ensuring ease 
of content management on the server.


### Usage Example:

```bash
$ go run main.go -port 8080 -dir /path/to/artifacts -https -cert /path/to/cert.pem -key /path/to/key.pem
```
In this example, the server will run on port 8080 and serve files and directories from the /path/to/artifacts directory. Additionally, the server will support HTTPS using the certificate and private key located at /path/to/cert.pem and /path/to/key.pem, respectively.

GoHttpFileServer provides a robust and efficient storage solution for distributing artifacts, build outputs, releases, and other crucial files, offering easy access to development teams during CI/CD processes and project management.