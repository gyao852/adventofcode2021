import io
class Result:
    @staticmethod
    def reader(filename):
        data = []
        with open(filename, "r") as f:
            for line in f:
                data.append(line.splitlines()[0])
        return data
class day1(Result):
    input = Result.reader("/Users/gyao/Downloads/1.txt")
    data = [int(i) for i in input]
    # def depth_counter(data):
    #     depth_counter = 0
    #     for i in range(1,len(data)):
    #         if data[i] > data[i-1]:
    #             depth_counter += 1
    #     return depth_counter
    def depth_counter(data):
        result = 0
        last_depth_window = data[0] + data[1] + data[2]
        for i in range(1,len(data)-2):
            depth_window = data[i] + data[i+1] + data[i+2]
            if depth_window > last_depth_window:
                result += 1
            last_depth_window = depth_window
        return result
    result = depth_counter(data)
    print(result)
