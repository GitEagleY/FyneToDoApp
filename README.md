<p align="center">
  <h1 align="center">FyneToDoApp</h1>
</p>

## Dependencies

#### This Fyne app relies on several dependencies that need to be installed before building and running the application. Make sure you have the following dependencies installed:

  - Go (version 1.15 or later): You can download and install Go from the official website: https://golang.org/

  - Fyne: Fyne is a Golang framework for building graphical user interfaces. You can find more information and installation instructions at the official Fyne documentation: [Fyne - Getting Started](https://developer.fyne.io/started/)

#### Please make sure all the dependencies are properly installed and configured before proceeding to build the application.

## Build Instructions

To build the app, follow these steps:

   1. Clone or download the source code for your Fyne app from the GitHub repository.

   2. Open a terminal or command prompt and navigate to the project directory.

   3. Run the following command to download and install any required Go dependencies:

    go mod download

  Now, you can build the application using the following command:

    fyne package -os <operating_system> -icon <icon_path> -name <app_name>

Replace <operating_system> with the target operating system for which you want to build the app (e.g., linux, windows, or macos).

Replace <icon_path> with the path to the application icon file (optional).

Replace <app_name> with the desired name for your application (optional).

This command will create an executable file for the specified operating system in the current directory.

## Troubleshooting

If you encounter any issues during the build process or while running the application, refer to the official Fyne documentation for troubleshooting information: [Fyne - Getting Started - Packaging](https://developer.fyne.io/started/packaging)

## License

This project is licensed under the BSD 3-Clause License.
