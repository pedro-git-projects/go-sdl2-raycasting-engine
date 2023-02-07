# Raycasting Engine using Go and SDL2

#### Table of Contents
  * [What is Raycasting](#what-is-raycasting)
  * [What is SDL](#what-is-sdl)
  * [Project Structure](#project-structure)
  * [Installing and Running](#installing-and-running)
  * [Documentation](#documentation)
  * [Todo](#todo)

## What is Raycasting

 The idea is as follows: the map is a 2D square grid, and each square will store a value which will be used to represent the texture that should be rendered in that square.

Then for every x-coordinate a ray will be cast, starting at the player position and with a direction that will depend on both, the player's rotation angle, and the screen's x-coordinate. Such a ray will move forward on the 2D map, until it collides with a solid coordinate. 

When this happens, the from that point to the player must be calculated and used determine how high the projected coordinate must be rendered to the screen. That implies that the further away the object, the smaller it is on screen.Conversely, the closer it is, the bigger it will  appear to be. 

 Although this are all 2D calculations, the final projection will appear to be 3D.

## What is SDL

Simple DirectMedia Layer is a cross-platform development library designed to provide low level access to audio, keyboard, mouse, joystick, and graphics hardware via OpenGL and Direct3D. It is used by video playback software, emulators, and popular games including Valve's award winning catalog and many Humble Bundle games. 

Here Go bindings were used to call the SDL library from Go source code.


## Project Structure

| Folder         | Purpose                                                                                                     |
|----------------|-------------------------------------------------------------------------------------------------------------|
| cmd            | stores application specific files and folders                                                               |
| cmd/root       | stores the app struct as well as app receiver functions, which constitute most of the application logic     |
| cmd/color      | stores the color buffer type, constructor, accessors and receivers                                          |
| cmd/game       | stores the Game struct, including the map and rays slice                                                    |
| cmd/player     | stores the turn and walk direction enums as well as the Player struct and associated functions              |
| cmd/ray        | stores the Ray struct, associated functions                                                                 |
| cmd/timekeeper | stores all application data relating to framerate, frame time and update pace                               |
| cmd/utils      | stores helper functions such as for calculating the distance between 2D coordinates and angle normalization |
| cmd/window     | stores constants and variables relating to the window dimensions, and number of rays                        |

## Installing and Running 

If you're going to build the program from source you'll need to install the following dependencies:

- [Go v1.13+](https://go.dev/dl/)
- [SDL2](https://github.com/libsdl-org/SDL/releases)
- [GNU Make(optional)](https://www.gnu.org/software/make/#download)

Once you have this requirements fullfilled it is enough to run:

```
$ make install 
```

Then

```
$ make run
```

Else you can download a pre compiled binary from this repository releases

## Documentation

To see the documentation run

```
$ make doc
```

## Todo

- Add textures

- Compile linux, windows and mac releases
