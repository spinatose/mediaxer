-- First draft of Media Orgnaizer (Mediaxer) version 1.0 --
Mediaxer first version is a Golang command line program that 
will take a folder as input and will organize all media files 
within the folder according to meta data each file into folders
that correspond to the date the file was created. 

First version will be very simple and will take one argument as
a command line argument- the folder in which media files reside.
The app's first iteration will not distinguish between file types.
It will just take all files in the folder and sort them according
to date created, create a subfolder for the date in format 20010911.

Example:
>mediaxer '/users/scot/tmp'

Where /users/scot/tmp is a folder that contains files:
tmp1.txt    -DateCreated - 2023-03-19

After execution:
/users/scot/tmp contains subfolder
20230319 which contains tmp1.txt


