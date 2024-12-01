import React, { useState, useEffect, useCallback } from "react";
import { useParams, Link } from "react-router-dom";
import axios from "axios";
import "./Room.css"; // Import the CSS file

const Room = ({ token }) => {
  const { roomId } = useParams();
  const [messages, setMessages] = useState([]);
  const [message, setMessage] = useState("");
  const [rooms, setRooms] = useState([]);
  const [loadingMessages, setLoadingMessages] = useState(true);
  const [loadingRooms, setLoadingRooms] = useState(true);

  useEffect(() => {
    const fetchMessages = async () => {
      try {
        const response = await axios.get(
          `http://localhost:8080/messages?roomId=${roomId}`,
          {
            headers: { Authorization: `Bearer ${token}` },
          }
        );
        setMessages(response.data || []);
      } catch (error) {
        console.error("Error fetching messages:", error);
      } finally {
        setLoadingMessages(false);
      }
    };

    fetchMessages();
  }, [roomId, token]); // This useEffect triggers when roomId or token changes

  const fetchRooms = useCallback(async () => {
    try {
      const response = await axios.get(`http://localhost:8080/rooms`, {
        headers: { Authorization: `Bearer ${token}` },
      });
      setRooms(response.data.rooms || []);
    } catch (error) {
      console.error("Error fetching rooms:", error);
    } finally {
      setLoadingRooms(false);
    }
  }, [token]);

  useEffect(() => {
    fetchRooms();
  }, [fetchRooms]); // This useEffect triggers only once when the component mounts

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

  const handleAddRoom = async (newRoom) => {
    await axios.post(`http://localhost:8080/rooms`, newRoom, {
      headers: { Authorization: `Bearer ${token}` },
    });
    fetchRooms(); // Fetch rooms again when a new room is added
  };

  if (loadingMessages || loadingRooms) {
    return <div>Loading...</div>;
  }

  return (
    <div id="mainContent">
      <h1 id="roomTitle">Room: {roomId}</h1>
      <div id="contentContainer">
        <div id="chatContainer">
          <div id="messages">
            {messages.map((msg, index) => (
              <div key={index} className="message">
                <div className="username">{msg.username}</div>
                <div className="content">{msg.content}</div>
              </div>
            ))}
          </div>
          <div id="messageInputContainer">
            <input
              id="messageInput"
              type="text"
              value={message}
              onChange={(e) => setMessage(e.target.value)}
              placeholder="Type a message"
            />
            <button id="sendMessageButton" onClick={handleSendMessage}>
              Send <i className="fas fa-paper-plane"></i>
            </button>
          </div>
        </div>
        <div id="sidebar">
          <h2>Rooms</h2>
          <ul id="roomList">
            {rooms.map((room) => (
              <li key={room.id}>
                <Link to={`/room/${room.id}`}>{room.name}</Link>
              </li>
            ))}
          </ul>
        </div>
      </div>
    </div>
  );
};

export default Room;
