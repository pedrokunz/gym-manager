export const fetchClasses = async () => {
  const query = `
    query {
      classes {
        id
        name
        trainer
        schedule
      }
    }
  `;

  const response = await fetch('http://localhost:8080/graphql', {
    method: 'POST',
    headers: { 'Content-Type': 'application/json' },
    body: JSON.stringify({ query }),
  });

  const result = await response.json();
  return result.data.classes;
};
