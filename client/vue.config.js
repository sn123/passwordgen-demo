module.exports = {
    devServer: {
        proxy: {
            '/api': {
                logLevel: 'debug',
                target: 'http://localhost:8080'
            }
        }
    }
};