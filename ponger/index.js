import protobuf from "protobufjs";
import http from "http";

const schema = await protobuf.load("../tick.proto");
const Tick = schema.lookupType("Tick");

http
  .createServer((req, res) => {
    const data = [];
    req
      .on("data", (chunk) => {
        data.push(chunk);
      })
      .on("close", () => {
        const buffer = Buffer.concat(data);
        const tick = Tick.decode(buffer);

        const time = tick.time;
        const date = new Date(
          time.seconds * 1e3 + Math.round(time.nanos / 1e6)
        );
        console.log(`pong ${tick.data} at ${date.toLocaleString()}`);
      });
    res.end();
  })
  .listen(8080);
