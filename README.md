# Raycasting Engine using Go and SDL2

#### Table of Contents
  * [What is Raycasting](#what-is-raycasting)
  * [What is SDL](#what-is-sdl)
  * [Implementation](#implementation)
  * [Project Structure](#project-structure)
  * [Installing and Running](#installing-and-running)
  * [Release](#release)
  * [Todo](#todo)

## What is Raycasting

 The idea is as follows: the map is a 2D square grid, and each square will store a value which will be used to represent the texture that should be rendered in that square.

Then for every x-coordinate a ray will be cast, starting at the player position and with a direction that will depend on both, the player's rotation angle, and the screen's x-coordinate. Such a ray will move forward on the 2D map, until it collides with a solid coordinate. 

When this happens, the from that point to the player must be calculated and used determine how high the projected coordinate must be rendered to the screen. That implies that the further away the object, the smaller it is on screen.Conversely, the closer it is, the bigger it will  appear to be. 

 Although this are all 2D calculations, the final projection will appear to be 3D.

## What is SDL

Simple DirectMedia Layer is a cross-platform development library designed to provide low level access to audio, keyboard, mouse, joystick, and graphics hardware via OpenGL and Direct3D. It is used by video playback software, emulators, and popular games including Valve's award winning catalog and many Humble Bundle games. 

Here Go bindings were used to call the SDL library from Go source code.

## Implementation


## Project Structure


## Installing and Running 

## Release

## Todo
