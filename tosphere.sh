#!/bin/sh

convert $1 $2

w=$(identify -format '%w' $1)
h=$(identify -format '%h' $1)

exiftool \
-ProjectionType=equirectangular \
-UsePanoramaViewer=True \
-PoseHeadingDegrees=0 \
-CroppedAreaImageWidthPixels=$w \
-CroppedAreaImageHeightPixels=$h \
-FullPanoWidthPixels=$w \
-FullPanoHeightPixels=$h \
-CroppedAreaLeftPixels=0 \
-CroppedAreaTopPixels=0 \
-overwrite_original \
$2
