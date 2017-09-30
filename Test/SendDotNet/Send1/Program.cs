using System;
using RabbitMQ.Client;
using System.Text;

namespace Send1
{
    class Program
    {
        static void Main(string[] args)
        {
            var factory = new ConnectionFactory() { HostName = "localhost", Port = 5672 };
            using (var connection = factory.CreateConnection())
            using (var channel = connection.CreateModel())
            {
                channel.QueueDeclare(queue: "canal1", durable: false, exclusive: false, autoDelete: false, arguments: null);
                string message = "";
                do
                {
                    Console.Write("Mensaje: ");
                    message = "canal1: ";
                    message += Console.ReadLine();

                    var body = Encoding.UTF8.GetBytes(message);
                    channel.BasicPublish(exchange: "", routingKey: "canal1", basicProperties: null, body: body);
                    Console.WriteLine(" [x] Sent {0}", message);
                } while (message != "salir");
            }

            Console.WriteLine(" Press [enter] to exit.");
            Console.ReadLine();
        }
    }
}
