# gorazor Example

Building the views

    cd example
    gorazor views views

## Rebuilding views automatically

Requires [node.js](http://nodejs.org) for auto building.

Install gulp and deps

    # inside of example dir
    npm install gulp -g
    npm install

Building the views with gulp

    gulp views

Watch for `views/**/*.gohtml`changes and rebuild

    gulp watch
