// src/components/Lobby.js
import React, { useState, useEffect } from "react";
import axios from "axios";

const Lobby = ({ token }) => {
  const [rooms, setRooms] = useState([]);
  const [loading, setLoading] = useState(true);

  useEffect(() => {
    const fetchRooms = async () => {
      try {
        const response = await axios.get("http://localhost:8080/rooms", {
          headers: { Authorization: `Bearer ${token}` },
        });
        setRooms(response.data.rooms);
      } catch (error) {
        console.error("Error fetching rooms:", error);
      } finally {
        setLoading(false);
      }
    };
    fetchRooms();
  }, [token]);

  if (loading) {
    return <div>Loading...</div>;
  }

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
