# Sort by start times
# keep track of previous end times
# iterate through meetings
# if current meeting starts before last end time, return false
# if make it all the way through, return true
def can_attend_meetings(meetings):
    sorted_meetings = sorted(meetings, key=lambda m: m[0])
    end_times = []

    for index, (start, end) in enumerate(sorted_meetings):
        if len(end_times):
            last_end_time = end_times[index - 1]

            if start < last_end_time:
                return False

        end_times.append(end)

    return True


intervals = [(10, 12), (6, 9), (13, 15)]
print(can_attend_meetings(intervals))










# [6,9][10,12][13,15]: True
# --6**9,10-11,12,13*15--------

# [6,9][8,12][13,15]: False
# --X**X----------------------- 6-9
# ----X***X-------------------- 8-12 (startTime is before last one's endTime)
# ---------X-X----------------- 13-15


# sort by end time:
# [8,12],[10,11],[16,17]
# 10-11, 8-12 16-17
# ----XX------------------------ 10-11
# --X***X---------------------- 8-12
# ---------XX------------------- 16-17

# end sorted
# [6,9][8,12][13,15]: False
# --X**X----------------------- 6-9
# ----X***X-------------------- 8-12 (startTime is before last one's endTime)
# ---------X-X----------------- 13-15

# 1-5, 2-4, 5-7
# 2-4
# 1-5
# 5-7
