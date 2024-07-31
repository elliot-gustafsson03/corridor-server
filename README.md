# Corridor Server

A server written in [Go](https://go.dev/) meant to be hosted on a local network.

## Usage

This server application is used for uploading and displaying images in a slideshow fashion, perfect for a common area, such as a [student corridor](https://www.afbostader.se/studentkorridor).

## Running the program

To use this application as intented, you will need an OS with a web browser and a Go compiler. Furthermore, the server will need to be connected to a local network.

### Starting the server

After cloning this repository, navigate into its directory and run the command

```bash
$ go run .
```

which will host the server on the local network on port 3333.

### Viewing the slideshow

On the same device, open up a web browser and navigate to [http://localhost:3333](http://localhost:3333). This is where images will appear.

### Uploading an image

On any device in the network, open a browser and navigate to [http://{ip-to-server}:3333/upload/](http://{ip-to-server}:3333/upload/). On this page you will be able to select an image and give it a description. After clicking upload, this information will be posted to the server and soon appear on the screen.

### Deleting an image

To delete an image, instead navigate to [http://{ip-to-server}:3333/delete/](http://{ip-to-server}:3333/delete/) and choose from the list which image to delete.
