# Corridor Server

A server written in [Go](https://go.dev/) meant to be hosted on a local network.

## Usage

This server application is used for hosting apps, (e.g. a slideshow displaying uploaded images), perfect for a common area, such as a [student corridor](https://www.afbostader.se/studentkorridor).

## Apps

At the moment, the supported apps are:

-   Slideshow
    > Users can upload images which will appear on the server screen.
-   Time
    > The current time and date will appear on screen, as well as an analog style clock and the current Swedish name day. (To change which names are displayed, edit the file /api/resources/namedays.json).

## Running the program

To use this application as intented, you will need an OS with a web browser and a Go compiler. Furthermore, the server will need to be connected to a local network.

### Starting the server

After cloning this repository, navigate into its directory and run the command

```bash
$ sudo go run .
```

which will host the server on the local network on port 80.

> Elevated privileges are required when hosting on the default port (:80)

### Viewing the slideshow

On the server device, open up a web browser and navigate to [http://localhost/apps/slideshow](http://localhost/apps/slideshow). This opens up the page where images will appear.

### Controlling the server

The server's behaviour can be controlled using any device on the local network. In a web browser, navigate to [http://{ip-to-server}](http://{ip-to-server}). Alternatively, use the host name of the server computer.
