// src/components/Lobby.js
import React, { useState, useEffect } from "react";
import axios from "axios";

const Lobby = ({ token }) => {
  const [rooms, setRooms] = useState([]);

  useEffect(() => {
    const fetchRooms = async () => {
      const response = await axios.get("http://localhost:8080/rooms", {
        headers: { Authorization: `Bearer ${token}` },
      });
      setRooms(response.data.rooms);
    };
    fetchRooms();
  }, [token]);

  return (
    <div>
      <h1>Lobby</h1>
      <ul>
        {rooms.map((room) => (
          <li key={room.id}>
            <a href={`/room?roomId=${room.id}`}>{room.name}</a>
          </li>
        ))}
      </ul>
    </div>
  );
};

export default Lobby;
