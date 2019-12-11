import React from "react";

export const Contact = ({ name, phone, address, favorites }) => {
    return (
        <div className="pa2 mb3 striped--near-white">
            <header className="b mb2">{name}</header>
            <div className="pl2">
                <p className="mb2">{phone}</p>
                <p className="pre mb3">{address}</p>
                <p className="mb2">
                    <span className="fw5">Favorite colors: </span>
                    {favorites.colors.join(", ")}
                </p>
            </div>
        </div>
    );
};
