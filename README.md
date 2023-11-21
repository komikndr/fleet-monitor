# Fleet Monitor

## What is it?
Fleet monitoring is a passion project to *only* monitor drone/s. It is currently in the Alpha version.
It is built upon Wails to be a desktop application, however from the start, it is already being decoupled as API, and the
FrontEnd side so it can be modified into working in web or just API alone. 

## Connection to the drone
It works by MavLink protocol produced and sent by serial communication of 3DR(XBEE) telemetry. 
*IT DOES NOT* control the drone, it is just monitoring software. If you want to control the drone as a ground control
use QCGroundControl or other similar products

## Front End 
Using vite, react, shading, tailwind, and radix/s. It is currently in the Alpha version, since it only communicates using 

## Back End
Using Wails Go, Gin, and GORM