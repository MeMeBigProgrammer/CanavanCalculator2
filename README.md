# The Canavan Calculator 2
Hitting Statistics for Baseball Players

I made a program in 7th grade that was horrible, I've included that in `/assets`. You can decompile the `.jar` file and laugh... or cry.

I remade it by using the existing functionality in Excel and a small Golang program.

The program is named after the MICDS Varsity Baseball coach, [Mr. Canavan](https://www.linkedin.com/in/tim-canavan-0a88a483)

## What does this do?

The Excel sheet is structured with one `Calculations` sheet, in which all the statistics are calculated and displayed. All other sheets are raw At-Bat statistics, and you have one sheet for each player.

By Seeing what outcomes (Strikeout, Hit, Walk, HBP) happen most often in specific circumstances (Pitch Location, Pitch Type), one can gain the following information:

- How many pitches am I seeing per at bat?
- What counts am I working myself into?
- What pitch type(s) am I putting into play?
- Am I putting strikes into play?
- Is there a pitch/location I need to swing at more/less often?
- Am I consistently hitting the ball hard?
- Do I need to make an adjustment against a certain pitcher (or pitching staff)?

In the `Calculations` sheet, you can change the sheet name at the top, to switch whose data is being calculated easily.

The Go program can load the Excel sheet to make entering data easier and checking for invalid inputs (i.e. 4 strikes, hitting a Line Drive while walking)

## Known Issues
- Because the UI package lacks a way to display images, I embedded the image into the spreadsheet and use open-golang to show the image

## LICENSE
License is MIT

Thanks to the following projects that this relies on:
- For UI: https://github.com/andlabs/ui
- For Excel: https://github.com/360EntSecGroup-Skylar/excelize
- For Opening: https://github.com/skratchdot/open-golang
