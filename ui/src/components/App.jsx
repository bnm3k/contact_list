import React, { useState } from "react";
import { Contact } from "./Contact";

const testData = [
    {
        id: 1,
        name: "Mrs. Hettie Bergstrom I",
        phone: "+5009890369775",
        address: "52434 Arturo Trace Suite 414 West Thea, NV 96425-7487",
        favorites: {
            colors: ["Chartreuse", "PeachPuff", "DeepPink"]
        },
        created_at: Date.now(),
        updated_at: Date.now()
    },
    {
        id: 2,
        name: "Mrs. Catherine Emard II",
        phone: "+1373890369775",
        address: "74477 Mohamed Divide Meredithchester, NC 26546-4810",
        favorites: {
            colors: ["Tan", "DarkCyan", "ForestGreen"]
        },
        created_at: Date.now(),
        updated_at: Date.now()
    }
];

export const App = () => {
    let [contacts, setContact] = useState(testData);
    return (
        <div className="mw6 center pa3 sans-serif">
            <h1 className="mb4">Contacts</h1>
            {contacts.map(contactDetails => (
                <Contact key={contactDetails.id} {...contactDetails} />
            ))}
        </div>
    );
};
