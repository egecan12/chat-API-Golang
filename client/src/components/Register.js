// src/components/Register.js
import React, { useState } from "react";
import axios from "axios";
import "./Register.css"; // Create a separate CSS file for styles

const Register = () => {
  const [username, setUsername] = useState("");
  const [password, setPassword] = useState("");

  const handleSubmit = async (e) => {
    e.preventDefault();
    try {
      const response = await axios.post("http://localhost:8080/register", {
        username,
        password,
      });
      if (response.status === 200) {
        alert("Registration successful! Please log in.");
        window.location.href = "/login";
      }
    } catch (error) {
      alert("Registration failed: " + error.response.data.error);
    }
  };

  return (
    <div className="register-container">
      <h1>Register</h1>
      <form onSubmit={handleSubmit}>
        <input
          type="text"
          value={username}
          onChange={(e) => setUsername(e.target.value)}
          placeholder="Enter your username"
          required
        />
        <input
          type="password"
          value={password}
          onChange={(e) => setPassword(e.target.value)}
          placeholder="Enter your password"
          required
        />
        <button type="submit">Register</button>
        <a href="/login">Already have an account? Login</a>
      </form>
    </div>
  );
};

export default Register;
