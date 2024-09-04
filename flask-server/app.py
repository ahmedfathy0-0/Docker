from flask import Flask,jsonify
from flask_mysqldb import MySQL
import os

app = Flask(__name__)

app.config['MYSQL_HOST'] = os.getenv('DB_HOST', 'mysql')  
app.config['MYSQL_USER'] = os.getenv('DB_USER', 'myuser')  
app.config['MYSQL_PASSWORD'] = os.getenv('DB_PASSWORD', 'pass') 
app.config['MYSQL_DB'] = os.getenv('DB_NAME', 'greetings')  

mysql = MySQL(app)

@app.route('/health')
def Health():
    return "Hello, there!"

@app.route('/', methods=['GET'])
def greet():
    cur = mysql.connection.cursor()
    cur.execute('''SELECT * FROM greets''')
    data = cur.fetchall()
    cur.close()
    return jsonify(data)

if __name__ == '__main__':
    app.run(host='0.0.0.0',port=5000)

