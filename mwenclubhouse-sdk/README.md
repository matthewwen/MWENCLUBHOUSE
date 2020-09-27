# MWENCLUBHOUSE-SDK

Located at https://api.matthewwen.com. Hosted by Matthew Wen. Domain (DNS) from Google Domains. Encryption by Let's Encrypt.

This sdk is developed seperately from the actual application running on the server. This allows me to easily test the program through a program like Docker / GitLab. In other words, instead of running the application, then testing random API keys or parameters through GET, PUT, POST commands, I can run my test cases to test the methods directly.

As a result, I can code with "confidence" when using it inside mwenclubhouse-gateway.

## Files Inside this Repo
```
Algorithm: (Still Active) Data Structure Algorithms Used in SDK
ToDoPlanner: Program Written with AWS SDK to keep track of Tasks
```

## Algorithm / Algorithms
|Directory | Purpose | Used by | 
|---|---| ---|
|Heap| Sorting Using a Heap. Runtime is O(nlog(n)) and the space complexity is O(1).| Posts on Personal Website|

## Reason Abandoned
---------------------------
|Directory| Reason|
|-----| ---- |
|todoplanner| Todoist added Boards. I've been using it since freshman year, and even though I do look for other options, none make me as productive as Todoist. The Boards was the last thing I needed for the ideal ToDoPlanner. It would be better to create wrappers around Todoist to better enhance my productivity. 
---------------------------
