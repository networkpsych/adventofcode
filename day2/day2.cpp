#include <stdio.h>
#include <ctype.h>
#include <stdlib.h>
#include <string.h>
#include <strings.h>

#define MAX(a, b) (((a) > (b)) ? (a) : (b))
// define max because windows
/*
Using the examples from u/clbrri on reddit.

This solition was not the solution I used for day 2. 

controls: 12 red, 13 green, 14 blue

Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green
Game 2: 1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue
Game 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red
Game 4: 1 green, 3 red, 6 blue; 3 green, 6 red; 3 green, 15 blue, 14 red
Game 5: 6 red, 1 blue, 3 green; 2 blue, 1 red, 2 green

expect output: 8
*/

int main()
{
    FILE *fptr = fopen("example.txt", "r");
    char *f = (char *)malloc(256), c[8];
    int sum = 0;
    int game_total = 0;
    int id = 0;
    // puzzle has a hard set limit
    int max_r = 12;
    int max_b = 14;
    int max_g = 13;
    // using fscanf, you can set the format of the starting input
    while (fscanf(fptr, "Game %d:", &id) == 1) // fscanf(FILE = fptr, format = "Game %d", current index = &id)
    {
        int r, g, b, n;
        r = 0;
        g = 0;
        b = 0;
        int valid = 1;

        while (fscanf(fptr, "%d %s", &n, c) == 2) // set the format for our qualifying input
        {
            if (*c == 'r') // check if the 
            {
                r = MAX(r, n);
                if (r > max_r)
                {
                    valid = 0;
                }
            }
            else if (*c == 'g')
            {
                g = MAX(g, n);
                if (g > max_g)
                {
                    valid = 0;
                }
            }
            else if (*c == 'b')
            {
                b = MAX(b, n);
                if (b > max_b)
                {
                    valid = 0;
                }
            }
        }
        sum += r * g * b;
        if (valid == 1){
            game_total = game_total + id;
        }
    }
    fclose(fptr);
    printf("SUM=%1u\n", sum);
    printf("VALID GAMES=%1u\n", game_total);
}