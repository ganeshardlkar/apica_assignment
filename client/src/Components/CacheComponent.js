import React, { useState } from "react";
import axios from "axios";

function CacheComponent() {
  const [key, setKey] = useState("");
  const [value, setValue] = useState("");
  const [setMessage, setSetMessage] = useState("");

  const [keyToFind, setKeyToFind] = useState("");
  const [getResponse, setGetResponse] = useState("");

  const handleGet = async () => {
    try {
      const response = await axios.get(
        `http://localhost:8080/get?key=${keyToFind}`
      );
      console.log(response);
      setGetResponse(response.data);
    } catch (error) {
      console.log(error);
      setGetResponse(error.response.data);
    }
  };

  const handleSet = async () => {
    try {
      await axios.post(`http://localhost:8080/set?key=${key}&value=${value}`);
      setSetMessage("Value successfully set in cache.");
    } catch (error) {
      console.error("Error setting data:", error);
    }
  };

  return (
    <div>
      <h1>LRU CACHE</h1>
      <div>
        <div>
          <h3> Set key and value for Cache</h3>
          <div>
            <input
              onChange={(e) => setKey(e.target.value)}
              placeholder="Enter key"
            />
          </div>
          <div>
            <input
              onChange={(e) => setValue(e.target.value)}
              placeholder="Enter value"
            />
          </div>
          <button onClick={handleSet}>Set value</button>
          {setMessage && <p>Value: {setMessage}</p>}
        </div>

        <div>
          <h3>Get value from Cache</h3>
          <div>
            <input
              onChange={(e) => setKeyToFind(e.target.value)}
              placeholder="Enter key"
            />
          </div>
          <button onClick={handleGet}>Get value</button>
          {getResponse && <p>Value: {getResponse}</p>}
        </div>
      </div>
    </div>
  );
}

export default CacheComponent;
