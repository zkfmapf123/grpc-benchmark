import express from 'express'
import { v4 as uuid } from 'uuid'
import axios from 'axios'

const [
    PORT,
    SUB_HOST
] = [
        process.env.PORT,
        process.env.SUB_HOST
    ]

const app = express()

app.get("/", (req, res) => res.status(200).send("hello world HTTP"))
app.get("/health", (req, res) => res.status(200).send("success"))
app.get("/pub", async (req, res) => {

    const data = {
        name: uuid()
    }

    let response
    try {
        response = await axios.post(`http://${SUB_HOST}/sub`, {
            data
        })
    } catch (e) {
        console.log("err")
        return res.status(500).json(JSON.stringify(response, null, 4))
    }

    return res.status(200).send("success publish")
})

app.listen(PORT, () => {
    console.log(`Connect To HTTP Protocol:${PORT}`)
})

