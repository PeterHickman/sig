# sig - Get the first 16 bytes of a file and display it as hex. Similar to md5sum etc

This is a very specific itch. I upgraded my Mac M1 mini to an M4 mini using the Apple supplied utility by restoring a backup from Timemachine. I have been doing this migration for years without issue. Until this time

A few files were corrupted. Looking at them showed that they contained the same bytes repeated until the origional file size was reached. Odd indeed!

So I wrote this to get the first 16 bytes of each file so I could compare the signature of a known corrupt file to all the other files that were transfered and find the other corrupted files. Found just over 600 out of more than 1,708,579. The files were ok on my old M1 so I just copied them back

Never found the cause

Maybe someday this will me useful again or perhaps it will be a foundation to build something else
