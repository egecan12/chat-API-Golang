import React, { useState, useEffect } from "react";
import axios from "axios";
import "./Lobby.css";

const Lobby = ({ token }) => {
  const [rooms, setRooms] = useState([]);
  const [loading, setLoading] = useState(true);

  useEffect(() => {
    const fetchRooms = async () => {
      try {
        const response = await axios.get("http://localhost:8080/rooms", {
          headers: { Authorization: `Bearer ${token}` },
        });
        setRooms(response.data.rooms || []);
      } catch (error) {
        console.error("Error fetching rooms:", error);
      } finally {
        setLoading(false);
      }
    };
    fetchRooms();
  }, [token]);

  if (loading) {
    return <div className="lobby-loading">Loading...</div>;
  }

  return (
    <div className="lobby-container">
      <header className="lobby-header">
        <h1>Lobby</h1>
      </header>
      <div className="lobby-grid">
        {rooms.map((room) => (
          <a
            href={`/room/${room.id}`}
            key={room.id}
            className="lobby-room-button"
          >
            {room.name}
          </a>
        ))}
      </div>
    </div>
  );
};

export default Lobby;
