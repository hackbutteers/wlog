# wlog

* Set global formatting
* example: ("%Y-%m-%d %H:%M:%S.%e %l : %v");
* %v    The actual text to log  "some user text"
* %t    Thread id       "1232"
* %n    Logger's name   "some logger name"
* %l    The log level of the message    "debug", "info", etc
* %L    Short log level of the message  "D", "I", etc
* %a    Abbreviated weekday name        "Thu"
* %A    Full weekday name       "Thursday"
* %b    Abbreviated month name  "Aug"
* %B    Full month name "August"
* %c    Date and time representation    "Thu Aug 23 15:35:46 2014"
* %C    Year in 2 digits        "14"
* %Y    Year in 4 digits        "2014"
* %D    Same as %x
* %x    Short MM/DD/YY date     "08/23/14"
* %m    Month 1-12      "11"
* %d    Day of month 1-31       "29"
* %H    Hours in 24 format 0-23 "23"
* %I    Hours in 12 format 1-12 "11"
* %M    Minutes 0-59    "59"
* %S    Seconds 0-59    "58"
* %e    Millisecond part of the current second 0-999    "678"
* %f    Microsecond part of the current second 0-999999 "056789"
* %F    Nanosecond part of the current second 0-999999999       "256789123"
* %p    AM/PM   "AM"
* %r    12 hour clock   "02:55:02 pm"
* %R    24-hour HH:MM time, equivalent to %H:%M "23:55"
* %T    Same as %X
* %X    ISO 8601 time format (HH:MM:SS), equivalent to %H:%M:%S "23:55:59"
* %z    ISO 8601 offset from UTC in timezone ([+/-]HH:MM)       "+02:00"
* %%    The % sign      "%"
* %+    fmtlog's default format "[2014-31-10 23:46:59.678] [info] [mylogger] Some message"