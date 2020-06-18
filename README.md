# trainee-head

## About

**trainee-head** is a from scratch developed program which outputs the first 10 lines of a file you specify. 
Seen like this it's an imitation of the original head command line command.

## Parameters

| Parameter    | Usage | 
| ------------- |-------------|
| **-n**      | Define how many lines should be printed.

## Usage Examples

In the following you'll see a quick example of how to use the head command. Here is what our text file looks like:

in.txt:

                one
                two
                three
                four
                five
                six
                seven
                eight
                nine
                ten
                eleven
                twelve
                
To print the first 10 lines of in.txt run `./head.linux-amd64 < in.txt`

The output is:

                one
                two
                three
                four
                five
                six
                seven
                eight
                nine
                ten
                
If you want to choose more than one file run `./head.linux-amd64 in.txt in2.txt in3.txt ...`

To print e.g the first 5 lines use the -n parameter `./head.linux-amd64 -n 5 in.txt`
