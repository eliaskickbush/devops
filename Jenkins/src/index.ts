import express from "express";
import fs from "fs/promises";
interface Config {
    port: number;
    hostname: string;
}

async function fetchConfig(filename: string): Promise<Config> {
    const file = await fs.readFile(filename);
    const parsedFile: Config = JSON.parse(file.toString());
    return new Promise((resolve,_)=> {
        resolve(parsedFile);
    });
}

async function main() {
    let app = express();
    const config = await fetchConfig('./config.json');

    app.get("/ping", (req,res) => {
        console.log(`Request to /ping from: ` + req.ip + req.socket);
        res.status(200).send("pong");
    });

    app.use((_req,res,_next) => {
        res.status(404).send("not found");
    });

    app.listen(config.port, config.hostname, () => {
        console.log(`Listening on ${config.hostname}:${config.port}...`);
    });
}

main().then(_ => {

}).catch(err => {
    console.error("There was an error:" + err);
})