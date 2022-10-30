# ----------- STATESTORE --------------------
POST http://localhost:3000/store/key2
```
this is stored
```

GET http://localhost:3000/retrieve/key2

HTTP/* 200
[Asserts]
body contains "this is stored"

# --------------- PUBSUB -----------------------
