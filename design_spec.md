# Penguin Design Specs

## File : Image_Fetcher

###### Method : FetchImageFromExternalSource()
- Fetches Image from given (hard-coded for now) URL
- stores it with a random name in a given (hard-coded for now) temp directory
- returns a path to the file.

## File : Image_Tracker

###### Variable : ActiveImageQueue
- ordered list of currently valid images

###### Method : MoveImageToDirectory(src, dest)
- moves an image from temp directory to actual directory.

###### Method : AddImageToActiveList()
- adds an image to the ActiveImageQueue and removes earliest image.

###### Method : IsNewImage()
- checks if the image is new, i.e it is different from the last image in the queue.

## File : Periodic_Scheduler

###### Method : ScheduleAndRepeat(task, period)
- executes a given task at periodic intervals.



###### Sources

- 