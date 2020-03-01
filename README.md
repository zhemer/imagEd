# imagEd - simple JPEG image editor

imagEd is a simple image editor, that duplicates specified line or column of JPEG image to selected image's location and saves it into new file.

```console
$ ./imagEd -h
Simple image editor, that duplicates specified line or column of JPEG image to selected image's location and saves it into new file
Usage of ./imagEd:
  -column int
    	Column to copy (default -1)
  -columnEnd int
    	Column to copy end
  -columnStart int
    	Column to copy start
  -imgDst string
    	Image file for output
  -imgSrc string
    	Image file
  -line int
    	Line to copy (default -1)
  -lineEnd int
    	Line to copy end
  -lineStart int
    	Line to copy start

Example:
  ./imagEd -imgSrc plan9_glenda.jpg -line 377 -lineStart 385 -lineEnd 478 -column 250 -columnStart 330 -columnEnd 375
```

Given [this image](https://upload.wikimedia.org/wikipedia/commons/thumb/a/a5/Glenda_bunny_mascot_of_plan_9_from_bell_black.jpg/375px-Glenda_bunny_mascot_of_plan_9_from_bell_black.jpg) and command from example you will get the image:

| | | |
| ![plan9](plan9_glenda.jpg "Plan 9 Glenda") | ==>> | ![plan9](copy_plan9_glenda.jpg "Plan 9 Glenda") |

| ------ | ------ | ------ |
| cell | cell | cell |
