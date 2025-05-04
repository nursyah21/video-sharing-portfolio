ffprobe -v quiet -i myname-malgeum.mp4 -show_streams
ffmpeg -i myname-malgeum.mp4 -ss 00:00:05 -vframes 1 output.jpg
ffmpeg -v quiet -i myname-malgeum.mp4 -c copy -hls_list_size 0 -hls_time 4  video.m3u8
ffmpeg -v quiet -i myname-malgeum.mp4  -c:v libx264 -vf scale=640:360 -hls_list_size 0 -hls_time 4  360p.m3u8