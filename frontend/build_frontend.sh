#!/bin/bash

# This builds the frontend assets that will be compiled into
# minified JavaScript and CSS and embedded into the app.
# Under the hood, it uses "vite build". 
npm run build

# This copies some of our essential images into the frontend/dist
# directory, so they'll appear in the app. Vite copies some images
# that it finds in stylesheets, and it renames them all, which is
# really nice for breaking things.
#
# Vite does not handle images referenced in Go's HTML templates, so
# we have to do that ourselves. There's a way to do this in the vite
# config, but after reading their fucking docs for four fucking hours 
# and getting only errors, I decided one line of bash script would be
# a lot easier. So here it is.
mkdir -p dist/assets/src/img && cp src/img/*.png dist/assets/src/img
