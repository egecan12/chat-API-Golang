// src/components/Room.js
import React, { useState, useEffect } from "react";
import { useParams } from "react-router-dom";
import axios from "axios";

const Room = ({ token }) => {
  const { roomId } = useParams();
  const [messages, setMessages] = useState([]);
  const [message, setMessage] = useState("");

  useEffect(() => {
    const fetchMessages = async () => {
      const response = await axios.get(
        `http://localhost:8080/messages?roomId=${roomId}`,
        {
          headers: { Authorization: `Bearer ${token}` },
        }
      );
      setMessages(response.data);
    };
    fetchMessages();
  }, [roomId, token]);

  const handleSendMessage = async () => {
    const payload = JSON.parse(atob(token.split(".")[1]));
    const newMessage = {
      content: message,
      username: payload.username,
      userId: payload.user_id,
    };
    await axios.post(`http://localhost:8080/messages`, newMessage, {
      headers: { Authorization: `Bearer ${token}` },
    });
    setMessages([...messages, newMessage]);
    setMessage("");
  };

  return (
    <div>
      <h1>Room: {roomId}</h1>
      <div>
        {messages.map((msg, index) => (
          <div key={index}>
            <strong>{msg.username}</strong>: {msg.content}
          </div>
        ))}
      </div>
      <input
        type="text"
        value={message}
        onChange={(e) => setMessage(e.target.value)}
      />
      <button onClick={handleSendMessage}>Send</button>
    </div>
  );
};

export default Room;
