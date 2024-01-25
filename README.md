# Fibonacci sequence API

## Endpoints (all of them are GET requests):
**/previous** - returns the previous number of the Fibonacci sequence, the sequence is not changed.<br>
**/current** - returns the current number of the Fibonacci sequence, the sequence is not changed.<br>
**/next** - returns the next number of the Fibonacci sequence and adds a new number to the sequence.<br>
**/reset** - resets the sequence to the beginning.<br>
**/make_trouble** - causes a panic for testing reasons to check recovery.<br>


## Example:

http://127.0.0.1:8080/current returns **0**<br>
http://127.0.0.1:8080/next returns **1**<br>
http://127.0.0.1:8080/next returns **1**<br>
http://127.0.0.1:8080/next returns **2**<br>
http://127.0.0.1:8080/previous returns **1**<br>
http://127.0.0.1:8080/current returns **2**<br>
http://127.0.0.1:8080/reset and then http://127.0.0.1:8080/current returns **0**<br>