This CLI tool helps you find train tickets by checking availability not just on your selected route but also from nearby stations. If tickets are unavailable from your boarding station, it searches for seats from earlier stations where they might still be open. Similarly, it looks for tickets to stations beyond your destination to maximize the chances of finding an available seat.

By scraping data from Paytm and ConfirmTkt, the tool provides a comprehensive list of available trains and smart alternatives. This approach helps you book tickets strategically, ensuring you don’t miss out on a journey due to limited direct availability

### Usage 
```

go run main.go SRC DEST YYYYMMDD

```

- SRC : Source Station Code
- DEST : Destination Station Code
- YYYYMMDD : Journey Start Date
