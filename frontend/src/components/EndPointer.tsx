import React, { useState } from 'react'
import axios from 'axios'

const EndPointer = () => {

    const [isLoading, setIsLoading] = useState<boolean>(false)
    const [randomNumber, setRandomNumber] = useState<number>()


    const handleClick = () => {
        setIsLoading(true)
        axios.get("http://localhost:8080/")
            .then((response) => setRandomNumber(response.data.randomNumber))
            .catch((error) => console.log(error))
            .finally(() => setIsLoading(false))

    }
    return (
        <div>
            <button onClick={handleClick}>GET Request</button>
            {isLoading && <h1>Loading...</h1>}
            <h1>{randomNumber}</h1>


        </div>
    )
}

export default EndPointer
