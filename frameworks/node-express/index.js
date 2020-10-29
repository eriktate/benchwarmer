const express = require("express");

const host = process.env.BENCH_HOST
const port = process.env.BENCH_PORT

const app = express();
app.use(express.json());

app.get("/hello", (req, res) => {
  res.send("Hello, World!");
});

app.post("/json", (req, res) => {
  const { greeting, name } = req.body;
  res.send(JSON.stringify({
    msg: `${greeting} ${name}`
  }));
});

app.listen(port, host, () => {
  console.log(`Listening on ${host}:${port}`)
});
