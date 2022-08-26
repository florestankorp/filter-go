# filter-go

This is my first attempt at writing [Go](https://go.dev/).

When I set out to learn this language, I wanted to build something simple, fun and challenging:

This [problem-set](https://cs50.harvard.edu/x/2022/psets/4/filter/more/) felt like a worthwhile challenge because it had just the right amount of complexity paired with the playfulness of manipulating bytes, sometimes leading to unexpected and interesting glitches:

**Calculation errors when reversing the image**

![glitch-02](/assets/glitches/glitch-02.bmp)

**Calculation errors when computing averages**

![glitch-01](/assets/glitches/glitch-01.bmp)

# About 
This is a command line application that allows the user to input a BMP image and apply a transformation to it. The program then outputs the result.

The app can perform the following transformations:
*  Convert the image to grayscale
*  Reflect the image horizontally
*  Blur the image

# Requirements
You need Go to run this app. You can find instructions on how to install it [here](https://go.dev/doc/install)

# Usage
From the root of the project you can run 

`$ go build && ./filter-go [flag] <infile>.bmp <outfile>.bmp`

To try out the app use the sample BMP files stored in the `/assets` folder. Resulting files will be written to the `/out` folder.

## Flags

```sh
  -b    Blur image
  -g    Make image grayscale
  -r    Reflect image horizontally
```

# Thanks

All the [Discord Gophers](https://discord.com/channels/118456055842734083/118456055842734083) that took the time to answer my questions <3